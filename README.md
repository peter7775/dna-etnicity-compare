# DNA Etnicity Compare

## CLI application compare etnicity analyze of DNA from different web services. As input, you use CSV file with results from this services; the output shows the statistical result of the comparison made by the main methods used.
---
## Structure of CSV input
```
Service,Ethnicity,Percentage,Rating
MyHeritage,ethnic_group_1,48.5,0.9
MyHeritage,ethnic_group_2,29.4,0.9
MyHeritage,ethnic_group_3,20.0,0.9
MyHeritage,ethnic_group_4,15.8,0.9
FamilyTreeDNA,ethnic_group_1,58.0,0.3
FamilyTreeDNA,ethnic_group_2,29.0,0.3
FamilyTreeDNA,ethnic_group_3,6.0,0.3
FamilyTreeDNA,ethnic_group_4,12.0,0.3
Genomelink,ethnic_group_1,15.4,0.75
Genomelink,ethnic_group_2,37.8,0.75
Genomelink,ethnic_group_3,37.9,0.75
Genomelink,ethnic_group_4,15.7,0.75
LivingDNA,ethnic_group_1,15.4,0.2
LivingDNA,ethnic_group_2,0.1,0.2
LivingDNA,ethnic_group_3,25.3,0.2
LivingDNA,ethnic_group_4,7,0.2
```
## usage
``go run main.go compare etnicity.csv``

or use build from folder ```build/``` like
```./etcompare compare etnicity.csv ```



