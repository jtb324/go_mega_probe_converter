package main

import (
	"fmt"
)

func main() {
	fmt.Println("Initializing program to gather gene information from gnomAD...")
	//Parsing the inputs
	input_path, output_option, api_endpoint := parser()

	//going through the input file and getting a slice of all the genes
	gene_slice := parse_input(input_path)

	// getting a list of structs that have info about the gene of interest
	gene_info_list := fetch_response(api_endpoint, gene_slice, output_option)

	if output_option == "mysql" {
		// initializing the database
		db, dbName, _ := initialize_db()

		// making the table and insert values into it
		_ = make_table(db, dbName, gene_info_list)
	} else {
		write_to_file(gene_info_list, output_option)
	}

}
