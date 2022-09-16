# rate-my-plate
**Coding Problem:**\
A restaurant keeps a log of (eater_id, foodmenu_id) for all the diners. 
The eater_id is a unique number for every diner and foodmenu_id is unique for every food 
item served on the menu. 
Write a program that reads this log file and returns the top 3 menu items consumed. 
If you find a eater_id with the same foodmenu_id more than once then show an error.

Expected output:
1. Code that can handle the above problem statement\
2. Testcases (with example log files) that checks all the possible conditions\

```
git clone https://github.com/yasin-cs-ko-ak/rate-my-plate.git
cd rate-my-plate
go run rate-plate.go
```
Output:
```
➜  rate-my-plate go run ./rate-my-plate.go 
Successfully Opened log.json
----Top 3 FoodMenuIDs----
FoodMenuIDs: 236 Count: 5
FoodMenuIDs: 233 Count: 4
FoodMenuIDs: 235 Count: 3
Error EaterID: 001 has FoodMenuID: 233 more than once
Error EaterID: 002 has FoodMenuID: 234 more than once
Error EaterID: 003 has FoodMenuID: 236 more than once
```
