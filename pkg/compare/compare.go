	package compare

	import (
	    "encoding/csv"
	    "fmt"
	    "os"
	    "strconv"
	)

	type EthnicityResult struct {
	    Service    string
	    Ethnicity  string
	    Percentage float64
	    Rating     float64
	}

	func LoadCSV(filename string) ([]EthnicityResult, error) {
	    file, err := os.Open(filename)
	    if err != nil {
	        return nil, err
	    }
	    defer file.Close()

	    reader := csv.NewReader(file)
	    records, err := reader.ReadAll()
	    if err != nil {
	        return nil, err
	    }

	    var results []EthnicityResult
	    for i, record := range records {
	        if i == 0 {
	            continue
	        }
	        percentage, _ := strconv.ParseFloat(record[2], 64)
	        rating, _ := strconv.ParseFloat(record[3], 64)
	        results = append(results, EthnicityResult{
	            Service:    record[0],
	            Ethnicity:  record[1],
	            Percentage: percentage,
	            Rating:     rating,
	        })
	    }
	    return results, nil
	}





	func BayesianMethod(data []EthnicityResult) map[string]float64 {
		serviceData := make(map[string][]EthnicityResult)
		for _, result := range data {
			serviceData[result.Service] = append(serviceData[result.Service], result)
		}
	
		ethnicities := make(map[string]float64)
		counts := make(map[string]int)
		
		for _, result := range data {
			ethnicities[result.Ethnicity] = 0
			counts[result.Ethnicity]++
		}
	
		numEthnicities := float64(len(ethnicities))
		for ethnicity := range ethnicities {
			ethnicities[ethnicity] = 1.0 / numEthnicities
		}
	
		for _, results := range serviceData {
			rating := results[0].Rating 
	
			serviceProbabilities := make(map[string]float64)
			totalPercentage := 0.0
			
			for _, result := range results {
				totalPercentage += result.Percentage
			}
	
			for _, result := range results {
				normalizedPercentage := result.Percentage / totalPercentage
				serviceProbabilities[result.Ethnicity] = normalizedPercentage
			}
	
			for ethnicity := range ethnicities {
				if serviceProb, exists := serviceProbabilities[ethnicity]; exists {
					ethnicities[ethnicity] = ethnicities[ethnicity]*(1-rating) + serviceProb*rating
				}
			}
		}
	
		total := 0.0
		for _, prob := range ethnicities {
			total += prob
		}
		
		for ethnicity := range ethnicities {
			ethnicities[ethnicity] = (ethnicities[ethnicity] / total) 
		}
	
		return ethnicities
	}
	

	func AverageMethod(data []EthnicityResult) map[string]float64 {
	    averages := make(map[string]float64)
	    counts := make(map[string]int)

	    for _, result := range data {
	        averages[result.Ethnicity] += result.Percentage
	        counts[result.Ethnicity]++
	    }

	    for ethnicity := range averages {
	        averages[ethnicity] /= float64(counts[ethnicity])
	    }

	    return averages
	}

	func WeightedAverageMethod(data []EthnicityResult) map[string]float64 {
	    weightedAverages := make(map[string]float64)
	    totalWeights := make(map[string]float64)

	    for _, result := range data {
	        weightedAverages[result.Ethnicity] += result.Percentage * result.Rating
	        totalWeights[result.Ethnicity] += result.Rating
	    }

	    for ethnicity := range weightedAverages {
	        if totalWeights[ethnicity] > 0 {
	            weightedAverages[ethnicity] /= totalWeights[ethnicity]
	        }
	    }

	    return weightedAverages
	}

	func DisplayResults(bayesianResults, averageResults, weightedAverageResults map[string]float64) {
	    fmt.Println("\nVýsledky Bayesovské metody:")
	    for ethnicity, probability := range bayesianResults {
	        fmt.Printf("%s: %.2f%%\n", ethnicity, probability*100)
	    }

	    fmt.Println("\nVýsledky průměrování:")
	    for ethnicity, average := range averageResults {
	        fmt.Printf("%s: %.2f%%\n", ethnicity, average)
	    }

	    fmt.Println("\nVýsledky váženého průměrování:")
	    for ethnicity, weightedAverage := range weightedAverageResults {
	        fmt.Printf("%s: %.2f%%\n", ethnicity, weightedAverage)
	    }
	}

