package benchmark

import (
	"log"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	startTime := time.Now() // Start timing for individual record insertion
	user := User{Name: "Test user hipsiga for benchmark"}

	if err := db.Create(&user).Error; err != nil {
		b.Fatalf("============================ >>>>>>> Failed to insert record: %v", err)
	} else {
		elapsedTime := time.Since(startTime) // Calculate elapsed time for insertion
		log.Printf("Inserted record: ID=%d, Name=%s, Time elapsed=%s", user.ID, user.Name, elapsedTime)

	}
}

func BenchmarkMaxOpenCons1(b *testing.B) {
	dsn := "root:root@tcp(127.0.0.1:64867)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // if you need consistent data, set this to false
	})

	if err != nil {
		b.Fatalf("============================ >>>>>>> Failed to connect to PostgreSQL database: %v", err)
	}

	// check if the table exist
	if !db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if db.Migrator().DropTable((&User{})); err != nil {
			// handle error if the table does not exist
			b.Fatalf("============================ >>>>>>> Failed to drop User table: %v", err)
		}
	}

	// create the table if not exists
	db.AutoMigrate(&User{})
	// set pool

	sqlDB, err := db.DB()

	if err != nil {
		b.Fatalf("============================ >>>>>>> Failed to get database connection: %v", err)
		log.Fatalf("============================ >>>>>>> Failed to get database connection: %v", err)
	}

	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	startTime := time.Now() // Start timing
	totalRows := 0          // Initialize row counter

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
			totalRows++ // Increment row counter
		}
	})
	b.StopTimer()

	// Calculate and log execution time
	elapsedTime := time.Since(startTime)
	log.Printf("Benchmark completed: Total rows inserted=%d, Time elapsed=%s", totalRows, elapsedTime)
}

func BenchmarkMaxOpenCons10(b *testing.B) {
	dsn := "root:root@tcp(127.0.0.1:64867)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // if you need consistent data, set this to false
	})

	if err != nil {
		b.Fatalf("============================ >>>>>>> Failed to connect to PostgreSQL database: %v", err)
	}

	// check if the table exist
	if !db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if db.Migrator().DropTable((&User{})); err != nil {
			// handle error if the table does not exist
			b.Fatalf("============================ >>>>>>> Failed to drop User table: %v", err)
		}

		// if err := db.AutoMigrate(&User{}); err != nil {
		// 	b.Fatalf("============================ >>>>>>> Failed to migrate User table: %v", err)
		// }
	}

	// create the table if not exists
	db.AutoMigrate(&User{})
	// set pool

	sqlDB, err := db.DB()

	if err != nil {
		b.Fatalf("============================ >>>>>>> Failed to get database connection: %v", err)
		log.Fatalf("============================ >>>>>>> Failed to get database connection: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	startTime := time.Now() // Start timing
	totalRows := 0          // Initialize row counter

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
			totalRows++
		}
	})

	b.StopTimer()
	// Calculate and log execution time
	elapsedTime := time.Since(startTime)
	log.Printf("Benchmark completed: Total rows inserted=%d, Time elapsed=%s", totalRows, elapsedTime)
}
