package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	b "github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

type ConnectorHTTP struct {
	Backend b.Backend
	Host    string
	Port    uint16
	Router  *gin.Engine
}

func NewConnectorHTTP(backend b.Backend, host string, port uint16) *ConnectorHTTP {
	//gin.SetMode(gin.ReleaseMode)
	return &ConnectorHTTP{
		Backend: backend,
		Host:    host,
		Port:    port,
		Router:  gin.Default(),
	}
}

func serviceHandler(backend b.Backend) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req b.Request
		var resp b.Response
		if c.Bind(&req) == nil {
			err := backend.Service(&req, &resp)
			if err == nil {
				c.JSON(200, resp)
			} else {
				c.JSON(500, resp)
			}
		} else {
			c.JSON(400, resp)
		}
	}
}

func initializeHandler(backend b.Backend) func(c *gin.Context) {
	return func(c *gin.Context) {
		initialized := backend.Initialize()
		if initialized {
			c.Status(200)
		} else {
			c.Status(500)
		}
	}
}

func lookupHandler(backend b.Backend) func(c *gin.Context) {
	type resultLookupResultArray struct {
		Result []b.LookupResult `json:"result"`
	}
	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")
		zoneIds := c.Request.Header[http.CanonicalHeaderKey("X-RemoteBackend-zone-id")]
		zoneId := "1"
		if zoneIds != nil && len(zoneIds) >= 0 {
			zoneId = zoneIds[0]
		}
		lookupResultArray, err := backend.Lookup(qtype, qname, zoneId)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultLookupResultArray{
				Result: lookupResultArray,
			}
			c.JSON(200, responseBody)
		}
	}
}

func listHandler(backend b.Backend) func(c *gin.Context) {
	type resultListResultArray struct {
		Result []b.ListResult `json:"result"`
	}
	return func(c *gin.Context) {
		qname := c.Param("qname")
		domainId := c.Param("domain_id")
		zoneIds := c.Request.Header[http.CanonicalHeaderKey("X-RemoteBackend-zone-id")]
		zoneId := "1"
		if zoneIds != nil && len(zoneIds) >= 0 {
			zoneId = zoneIds[0]
		}
		listResultArray, err := backend.List(qname, domainId, zoneId)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultListResultArray{
				Result: listResultArray,
			}
			c.JSON(200, responseBody)
		}
	}
}

func getAllDomainsHandler(backend b.Backend) func(c *gin.Context) {
	type resultGetAllDomainsResultArray struct {
		Result []b.DomainInfoResult `json:"result"`
	}
	return func(c *gin.Context) {
		includeDisabled := c.Query("includeDisabled") == "true"
		domainInfoResultArray, err := backend.GetAllDomains(includeDisabled)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultGetAllDomainsResultArray{
				Result: domainInfoResultArray,
			}
			c.JSON(200, responseBody)
		}
	}
}

func getAllDomainMetadataHandler(backend b.Backend) func(c *gin.Context) {
	type resultLookupResultArray struct {
		Result map[string][]string `json:"result"`
	}
	return func(c *gin.Context) {
		qname := c.Param("qname")
		metadata, err := backend.GetAllDomainMetadata(qname)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultLookupResultArray{
				Result: metadata,
			}
			c.JSON(200, responseBody)
		}
	}
}

func getDomainMetadataHandler(backend b.Backend) func(c *gin.Context) {
	type resultDomainMetadataResultArray struct {
		Result []string `json:"result"`
	}
	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")
		metadata, err := backend.GetDomainMetadata(qname, qtype)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultDomainMetadataResultArray{
				Result: metadata,
			}
			c.JSON(200, responseBody)
		}
	}
}

func getDomainInfoHandler(backend b.Backend) func(c *gin.Context) {
	type resultDomainInfoResult struct {
		Result b.DomainInfoResult `json:"result"`
	}
	return func(c *gin.Context) {
		qname := c.Param("qname")
		domainInfoResult, err := backend.GetDomainInfo(qname)
		if err != nil {
			c.Status(500)
			c.Abort()
		} else {
			responseBody := resultDomainInfoResult{
				Result: domainInfoResult,
			}
			c.JSON(200, responseBody)
		}
	}
}

func (httpConnector *ConnectorHTTP) Config() (err error) {

	httpConnector.Router.POST("dnsapi/service", serviceHandler(httpConnector.Backend))
	httpConnector.Router.GET("dnsapi/initialize", initializeHandler(httpConnector.Backend))
	httpConnector.Router.GET("dnsapi/lookup/:qname/:qtype", lookupHandler(httpConnector.Backend))

	httpConnector.Router.GET("dnsapi/list/:domain_id/:qname", listHandler(httpConnector.Backend))

	httpConnector.Router.GET("dnsapi/getAllDomains", getAllDomainsHandler(httpConnector.Backend))
	httpConnector.Router.GET("dnsapi/getAllDomainMetadata/:qname", getAllDomainMetadataHandler(httpConnector.Backend))
	httpConnector.Router.GET("dnsapi/getDomainMetadata/:qname/:qtype", getDomainMetadataHandler(httpConnector.Backend))
	httpConnector.Router.GET("dnsapi/getDomainInfo/:qname", getDomainInfoHandler(httpConnector.Backend))
	return
}

func (httpConnector *ConnectorHTTP) Start() (err error) {
	addr := fmt.Sprintf("%s:%d", httpConnector.Host, httpConnector.Port)
	return httpConnector.Router.Run(addr)
}
