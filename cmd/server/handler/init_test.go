package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/buyer"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/carrier"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/employee"
	inboundOrders "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/locality"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/product"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productBatch"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/purchaseOrders"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/section"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/seller"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/warehouse"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Method that create server test
func CreateServer() *gin.Engine {

	// Method that initializes the database
	db, err := db.Init("test")
	if err != nil {
		fmt.Println(err)
	}

	repository := warehouse.NewRepository(db)
	service := warehouse.NewService(repository)
	warehouseHandler := NewWarehouse(*service)
	route := gin.Default()
	warehousesRoutes := route.Group("/api/v1/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	}

	productBatchRepository := productBatch.NewRepository(db)
	productBatchService := productBatch.NewService(productBatchRepository)
	productBatchHandler := NewProductBatch(*productBatchService)
	productBatchRoutes := route.Group("/api/v1/productBatches")
	{
		productBatchRoutes.POST("/", productBatchHandler.Store())
	}

	sellerRepository := seller.NewRepository(db)
	sellerService := seller.NewService(sellerRepository)
	sellerHandler := NewSeller(*sellerService)
	sellerRoutes := route.Group("/api/v1/sellers")
	{
		sellerRoutes.GET("/", sellerHandler.GetAll())
		sellerRoutes.GET("/:id", sellerHandler.Get())
		sellerRoutes.POST("/", sellerHandler.Store())
		sellerRoutes.PATCH("/:id", sellerHandler.Update())
		sellerRoutes.DELETE("/:id", sellerHandler.Delete())
	}
	localityRepository := locality.NewRepository(db)
	localityService := locality.NewService(localityRepository)
	localityHandler := NewLocality(*localityService)
	localityRoutes := route.Group("/api/v1/localities")
	{
		localityRoutes.GET("reportSellers/:id", localityHandler.GetReport())
		localityRoutes.GET("reportSellers/", localityHandler.GetGeneralReport())
		localityRoutes.POST("/", localityHandler.Store())
	}
	sectionRepository := section.NewRepository(db)
	sectionService := section.NewService(sectionRepository)
	sectionHandler := NewSection(*sectionService)
	sectionsRoutes := route.Group("/api/v1/sections")
	{
		sectionsRoutes.GET("/", sectionHandler.GetAll())
		sectionsRoutes.GET("/:id", sectionHandler.Get())
		sectionsRoutes.POST("/", sectionHandler.Store())
		sectionsRoutes.PATCH("/:id", sectionHandler.Update())
		sectionsRoutes.DELETE("/:id", sectionHandler.Delete())

		// Product Batches report endpoint
		sectionsRoutes.GET("/reportProducts/", productBatchHandler.ReportAll())
		sectionsRoutes.GET("/reportProducts", productBatchHandler.ReportBySection())
	}

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := NewProduct(*productService)
	productsRoutes := route.Group("/api/v1/products")
	{
		productsRoutes.GET("/", productHandler.GetAll())
		productsRoutes.POST("/", productHandler.Save())
		productsRoutes.GET("/:id", productHandler.Get())
		productsRoutes.PATCH("/:id", productHandler.Update())
		productsRoutes.DELETE("/:id", productHandler.Delete())
	}

	productRecordRepository := productRecord.NewRepository(db)
	productRecordService := productRecord.NewService(productRecordRepository)
	productRecordHandler := NewProductRecord(*productRecordService)
	productsRecordRoutes := route.Group("/api/v1/productsRecords")
	{

		productsRecordRoutes.POST("/", productRecordHandler.Save())
		productsRecordRoutes.GET("/", productRecordHandler.GetAll())
		productsRecordRoutes.GET("/:id", productRecordHandler.Get())
	}

	InboundOrdersRepository := inboundOrders.NewRepository(db)
	InboundOrdersService := inboundOrders.NewService(InboundOrdersRepository)
	InboundOrdersHandler := NewInboundOrders(*InboundOrdersService)
	InboundOrdersRoutes := route.Group("/api/v1/InboundOrders")
	{
		InboundOrdersRoutes.GET("/", InboundOrdersHandler.GetAll())
		InboundOrdersRoutes.GET("/:id", InboundOrdersHandler.Get())
		InboundOrdersRoutes.POST("/", InboundOrdersHandler.Store())
	}

	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := NewEmployee(*employeeService)
	employeeRoutes := route.Group("/api/v1/employees")
	{
		employeeRoutes.GET("/", employeeHandler.GetAll())
		employeeRoutes.GET("/:id", employeeHandler.Get())
		employeeRoutes.POST("/", employeeHandler.Store())
		employeeRoutes.PATCH("/:id", employeeHandler.Update())
		employeeRoutes.DELETE("/:id", employeeHandler.Delete())

		employeeRoutes.GET("/reportInboundOrders/", employeeHandler.GetInboundOrders())
		employeeRoutes.GET("/reportInboundOrders/:id", employeeHandler.GetInboundOrdersByEmployee())

	}
	buyerRepository := buyer.NewRepository(db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := NewBuyer(*buyerService)
	buyerRoutes := route.Group("/api/v1/buyers")
	{
		buyerRoutes.GET("/", buyerHandler.GetAll())
		buyerRoutes.GET("/:id", buyerHandler.Get())
		buyerRoutes.POST("/", buyerHandler.Store())
		buyerRoutes.PATCH("/:id", buyerHandler.Update())
		buyerRoutes.DELETE("/:id", buyerHandler.Delete())
		buyerRoutes.GET("/reportPurchaseOrders/", buyerHandler.GetPurchaseOrders())
		buyerRoutes.GET("/reportPurchaseOrders/:id", buyerHandler.GetPurchaseOrdersByBuyer())
	}

	carrierRepository := carrier.NewRepository(db)
	carrierService := carrier.NewService(carrierRepository)
	carrierHandler := NewCarrier(*carrierService)

	route.POST("/api/v1/carriers/", carrierHandler.Store())

	route.GET("/api/v1/localities/reportCarriers/", carrierHandler.GetAllReport())
	route.GET("/api/v1/localities/reportCarriers/:id/details", carrierHandler.GetReportDetails())
	route.GET("/api/v1/localities/reportCarriers/:id", carrierHandler.GetReport())

	purchaseOrderRepository := purchaseOrders.NewRepository(db)
	purchaseOrderService := purchaseOrders.NewService(purchaseOrderRepository)
	purchaseOrderHandler := NewPurchaseOrder(*purchaseOrderService)

	route.GET("/api/v1/PurchaseOrders/", purchaseOrderHandler.GetAll())
	route.GET("/api/v1/PurchaseOrders/:id", purchaseOrderHandler.Get())
	route.POST("/api/v1/PurchaseOrders/", purchaseOrderHandler.Store())

	return route
}

// Method that create Request test
func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	return req, httptest.NewRecorder()
}
