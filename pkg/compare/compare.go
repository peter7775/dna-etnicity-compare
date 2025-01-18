	package compare

	import (
	    "encoding/csv"
	    "fmt"
	    "math"
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


	func CalculateBayesian(results []EthnicityResult) map[string]float64 {
		ethnicities := make(map[string]bool)
		services := make(map[string]bool)
		for _, result := range results {
			ethnicities[result.Ethnicity] = true
			services[result.Service] = true
		}
	
		priors := make(map[string]float64)
		for ethnicity := range ethnicities {
			priors[ethnicity] = 1.0 / float64(len(ethnicities))
		}
	
		for service := range services {
			likelihoods := make(map[string]float64)
			totalLikelihood := 0.0
			serviceRating := 0.0
	
			for _, result := range results {
				if result.Service == service {
					serviceRating = result.Rating
					break
				}
			}
	
			for ethnicity := range ethnicities {
				likelihood := 0.0
				for _, result := range results {
					if result.Service == service && result.Ethnicity == ethnicity {
						likelihood = math.Log1p(result.Percentage) * serviceRating
						break
					}
				}
				if likelihood == 0 {
					likelihood = 0.01 * serviceRating
				}
				likelihoods[ethnicity] = likelihood
				totalLikelihood += likelihood
			}
	
			for ethnicity := range likelihoods {
				likelihoods[ethnicity] /= totalLikelihood
			}
	
			posteriors := make(map[string]float64)
			totalPosterior := 0.0
			for ethnicity := range ethnicities {
				posterior := likelihoods[ethnicity] * priors[ethnicity]
				posteriors[ethnicity] = posterior
				totalPosterior += posterior
			}
	
			for ethnicity := range posteriors {
				posteriors[ethnicity] /= totalPosterior
				priors[ethnicity] = posteriors[ethnicity]
			}
		}
	
		return priors
	}



	func BayesianMethod(data []EthnicityResult) map[string]float64 {
		ethnicities := make(map[string]bool)
		services := make(map[string]bool)
		for _, result := range data {
			ethnicities[result.Ethnicity] = true
			services[result.Service] = true
		}
	
		priors := make(map[string]float64)
		for ethnicity := range ethnicities {
			priors[ethnicity] = 1.0 / float64(len(ethnicities))
		}
	
		for service := range services {
			likelihoods := make(map[string]float64)
			totalLikelihood := 0.0
			serviceRating := 0.0
	
			for _, result := range data {
				if result.Service == service {
					serviceRating = result.Rating
					break
				}
			}
	
			for ethnicity := range ethnicities {
				likelihood := 0.0
				for _, result := range data {
					if result.Service == service && result.Ethnicity == ethnicity {
						// Use square root transformation instead of log
						likelihood = math.Sqrt(result.Percentage) * serviceRating
						break
					}
				}
				if likelihood == 0 {
					// Use a higher minimum likelihood
					likelihood = 0.1 * serviceRating
				}
				likelihoods[ethnicity] = likelihood
				totalLikelihood += likelihood
			}
	
			for ethnicity := range likelihoods {
				likelihoods[ethnicity] /= totalLikelihood
			}
	
			posteriors := make(map[string]float64)
			totalPosterior := 0.0
			for ethnicity := range ethnicities {
				// Apply service rating as weight
				posterior := (likelihoods[ethnicity]*serviceRating + priors[ethnicity]) / (serviceRating + 1)
				posteriors[ethnicity] = posterior
				totalPosterior += posterior
			}
	
			for ethnicity := range posteriors {
				posteriors[ethnicity] /= totalPosterior
				priors[ethnicity] = posteriors[ethnicity]
			}
		}
	
		return priors
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

