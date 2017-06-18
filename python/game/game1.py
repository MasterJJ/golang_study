print "Welcom!"

g = input( "Guess the number: ")
guess = int(g)
if guess == 5:
    print "You win!"
else:
    if guess > 5:
        print "Too hight"
    else:
        print "Too low"

print "Game Over"

