package main

import (
	"context"
	"github/felipecardosodeoliveira/golang-labs-auction/configuration/database/mongodb"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/api/web/controller/auction_controller"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/api/web/controller/bid_controller"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/api/web/controller/user_controller"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/database/auction"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/database/bid"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/infra/database/user"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/usecase/auction_usecase"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/usecase/bid_usecase"
	"github/felipecardosodeoliveira/golang-labs-auction/internal/usecase/user_usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionsController := initDependencies(databaseConnection)

	router.GET("/auction", auctionsController.FindAuctions)
	router.GET("/auction/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auction", auctionsController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(
		user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(
		auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}

// func main() {
// 	ctx := context.Background()

// 	if err := godotenv.Load("cmd/auction/.env"); err != nil {
// 		log.Fatal("Error trying to load env variables")
// 		return
// 	}

// 	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	router := gin.Default()

// 	userController, bidController, auctionsController := initDependencies(databaseConnection)

// 	router.GET("/auction", auctionsController.FindAuctions)
// 	router.GET("/auction/:auctionId", auctionsController.FindAuctionById)
// 	router.POST("/auction", auctionsController.CreateAuction)
// 	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
// 	router.POST("/bid", bidController.CreateBid)
// 	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
// 	router.GET("/user/:userId", userController.FindUserById)

// 	router.Run(":8080")
// }

// func initDependencies(database *mongo.Database) (
// 	userController *user_controller.UserController,
// 	bidController *bid_controller.BidController,
// 	auctionController *auction_controller.AuctionController) {

// 	auctionRepository := auction.NewAuctionRepository(database)
// 	bidRepository := bid.NewBidRepository(database, auctionRepository)
// 	userRepository := user.NewUserRepository(database)

// 	userController = user_controller.NewUserController(
// 		user_usecase.NewUserUseCase(userRepository))
// 	auctionController = auction_controller.NewAuctionController(
// 		auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
// 	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

// 	return
// }