package routes

import (
	"final_api/controllers" // Gantilah dengan path yang sesuai untuk controller Anda

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	server.GET("/", controllers.Home)

	// Rute untuk Users
	server.GET("/users", controllers.GetUsers)          // Mendapatkan semua pengguna
	server.POST("/users", controllers.SignUp)           // Membuat pengguna baru
	server.PUT("/users/:id", controllers.UpdateUser)    // Memperbarui pengguna berdasarkan ID
	server.DELETE("/users/:id", controllers.DeleteUser) // Menghapus pengguna berdasarkan ID

	// Rute untuk Transactions
	server.GET("/transactions", controllers.GetTransactions)          // Mendapatkan semua transaksi
	server.POST("/transactions", controllers.CreateTransaction)       // Membuat transaksi baru
	server.PUT("/transactions/:id", controllers.UpdateTransaction)    // Memperbarui transaksi berdasarkan ID
	server.DELETE("/transactions/:id", controllers.DeleteTransaction) // Menghapus transaksi berdasarkan ID
}
