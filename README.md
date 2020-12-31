# Random Luck
**Random Luck** Is an application used to create a random luck numbers based on quantity numbers to Mega Senna.
The numbers allowed to create the result is between 1 and 60 and we can only select a quantity of numbers between 6 and 9 for each game.

### How to use
In the terminal, you need to build the application using the command `go build` in the root folder.
After that you will run the binary in your terminal using the command `./random-luck`.



The application will ask you to fill number of games and there is no limit on values here, but need to be greater than zero.
It will also ask you to fill the numbers to game that have a limit between 6 and 9.

Now, the number will be shown for you in the output as we can see below:
```
** Please enter the numbers to generate the games **
Select number of games:
2
Select numbers to game:
6

** Generating **
Total games generated: 2

** Results: **
Numbers to game: 1
[3 14 27 45 48 53]

Numbers to game: 2
[1 14 17 35 44 59]

```

# Good Luck...