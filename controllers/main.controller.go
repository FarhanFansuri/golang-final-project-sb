package controllers

import (
	"final_api/models" // Gantilah dengan path yang sesuai untuk model Anda
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Inisialisasi DB
var DB *gorm.DB

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Home",
	})
}

// GetUsers akan mengembalikan daftar semua pengguna
func GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// GetTransactions akan mengembalikan daftar semua transaksi
func GetTransactions(ctx *gin.Context) {
	var transactions []models.Transaction
	if err := DB.Find(&transactions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// CreateUser untuk menambahkan user baru
func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// CreateTransaction untuk menambahkan transaksi baru
func CreateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := DB.Create(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}

// UpdateUser untuk memperbarui data pengguna berdasarkan ID
func UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind data yang baru dari request
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Update user
	if err := DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// UpdateTransaction untuk memperbarui transaksi berdasarkan ID
func UpdateTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("id")
	var transaction models.Transaction
	if err := DB.First(&transaction, transactionID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Bind data yang baru dari request
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Update transaction
	if err := DB.Save(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully", "transaction": transaction})
}

// DeleteUser untuk menghapus user berdasarkan ID
func DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hapus user
	if err := DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// DeleteTransaction untuk menghapus transaksi berdasarkan ID
func DeleteTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("id")
	var transaction models.Transaction
	if err := DB.First(&transaction, transactionID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Hapus transaksi
	if err := DB.Delete(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
