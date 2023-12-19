package main

import "fmt"

type ComplexETL struct{}

func (etl *ComplexETL) Extract() {
	fmt.Println("Extracting...")
}
func (etl *ComplexETL) Transform() {
	fmt.Println("Transforming...")
}
func (etl *ComplexETL) Load() {
	fmt.Println("Loading...")
}

type ETLFacade struct {
	ETL *ComplexETL
}

func (etl *ETLFacade) Process() {
	etl.ETL.Extract()
	etl.ETL.Transform()
	etl.ETL.Load()
}

func main() {
	etl := &ETLFacade{&ComplexETL{}}
	etl.Process()
}
