package main

import (
	"fmt"
	"os"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/docs"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/purchaseOrders"

	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/carrier"
	inboundOrders "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/locality"

	product "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/product"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productBatch"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/section"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/seller"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/warehouse"
	db "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/pkg/database"

	buyer "github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/buyer"
	"github.com/extmatperez/meli_bootcamp10_sprints/tree/team3_sprint4/team3/sprint4/bootcamp-go/internal/employee"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API - TEAM 3
// @version 1.0
// @description This API Handle MELI Fresh Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	db, err := db.Init("production")
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	//Documentation
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	purchaseOrderRepository := purchaseOrders.NewRepository(db)
	purchaseOrderService := purchaseOrders.NewService(purchaseOrderRepository)
	purchaseOrderHandler := handler.NewPurchaseOrder(*purchaseOrderService)

	router.GET("/api/v1/PurchaseOrders/", purchaseOrderHandler.GetAll())
	router.GET("/api/v1/PurchaseOrders/:id", purchaseOrderHandler.Get())
	router.POST("/api/v1/PurchaseOrders/", purchaseOrderHandler.Store())

	carrierRepository := carrier.NewRepository(db)
	carrierService := carrier.NewService(carrierRepository)
	carrierHandler := handler.NewCarrier(*carrierService)
	router.POST("/api/v1/carriers/", carrierHandler.Store())

	router.GET("/api/v1/localities/reportCarriers/", carrierHandler.GetAllReport())
	router.GET("/api/v1/localities/reportCarriers/:id/details", carrierHandler.GetReportDetails())
	router.GET("/api/v1/localities/reportCarriers/:id", carrierHandler.GetReport())

	sellerRepository := seller.NewRepository(db)
	sellerService := seller.NewService(sellerRepository)
	sellerHandler := handler.NewSeller(*sellerService)
	sellerRoutes := router.Group("/api/v1/sellers")
	{
		sellerRoutes.GET("/", sellerHandler.GetAll())
		sellerRoutes.GET("/:id", sellerHandler.Get())
		sellerRoutes.POST("/", sellerHandler.Store())
		sellerRoutes.PATCH("/:id", sellerHandler.Update())
		sellerRoutes.DELETE("/:id", sellerHandler.Delete())
	}

	localityRepository := locality.NewRepository(db)
	localityService := locality.NewService(localityRepository)
	localityHandler := handler.NewLocality(*localityService)
	localityRoutes := router.Group("/api/v1/localities")
	{
		localityRoutes.GET("reportSellers/:id", localityHandler.GetReport())
		localityRoutes.GET("reportSellers/", localityHandler.GetGeneralReport())
		localityRoutes.POST("/", localityHandler.Store())
	}

	warehouseRepository := warehouse.NewRepository(db)
	warehouseService := warehouse.NewService(warehouseRepository)
	warehouseHandler := handler.NewWarehouse(*warehouseService)
	warehousesRoutes := router.Group("/api/v1/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	}

	productBatchRepository := productBatch.NewRepository(db)
	productBatchService := productBatch.NewService(productBatchRepository)
	productBatchHandler := handler.NewProductBatch(*productBatchService)
	productBatchRoutes := router.Group("/api/v1/productBatches")
	{
		productBatchRoutes.POST("/", productBatchHandler.Store())
	}

	sectionRepository := section.NewRepository(db)
	sectionService := section.NewService(sectionRepository)
	sectionHandler := handler.NewSection(*sectionService)
	sectionsRoutes := router.Group("/api/v1/sections")
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
	productHandler := handler.NewProduct(*productService)
	productsRoutes := router.Group("/api/v1/products")
	{
		productsRoutes.GET("/", productHandler.GetAll())
		productsRoutes.POST("/", productHandler.Save())
		productsRoutes.GET("/:id", productHandler.Get())
		productsRoutes.PATCH("/:id", productHandler.Update())
		productsRoutes.DELETE("/:id", productHandler.Delete())
	}

	productRecordRepository := productRecord.NewRepository(db)
	productRecordService := productRecord.NewService(productRecordRepository)
	productRecordHandler := handler.NewProductRecord(*productRecordService)
	productsRecordRoutes := router.Group("/api/v1/productsRecords")
	{

		productsRecordRoutes.POST("/", productRecordHandler.Save())
		productsRecordRoutes.GET("/", productRecordHandler.GetAll())
		productsRecordRoutes.GET("/:id", productRecordHandler.Get())
	}

	InboundOrdersRepository := inboundOrders.NewRepository(db)
	InboundOrdersService := inboundOrders.NewService(InboundOrdersRepository)
	InboundOrdersHandler := handler.NewInboundOrders(*InboundOrdersService)
	InboundOrdersRoutes := router.Group("/api/v1/InboundOrders")
	{
		InboundOrdersRoutes.GET("/", InboundOrdersHandler.GetAll())
		InboundOrdersRoutes.GET("/:id", InboundOrdersHandler.Get())
		InboundOrdersRoutes.POST("/", InboundOrdersHandler.Store())
	}

	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployee(*employeeService)
	employeeRoutes := router.Group("/api/v1/employees")
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
	buyerHandler := handler.NewBuyer(*buyerService)
	buyerRoutes := router.Group("/api/v1/buyers")
	{
		buyerRoutes.GET("/", buyerHandler.GetAll())
		buyerRoutes.GET("/:id", buyerHandler.Get())
		buyerRoutes.POST("/", buyerHandler.Store())
		buyerRoutes.PATCH("/:id", buyerHandler.Update())
		buyerRoutes.DELETE("/:id", buyerHandler.Delete())
		buyerRoutes.GET("/reportPurchaseOrders/", buyerHandler.GetPurchaseOrders())
		buyerRoutes.GET("/reportPurchaseOrders/:id", buyerHandler.GetPurchaseOrdersByBuyer())

	}

	router.Run()

}
