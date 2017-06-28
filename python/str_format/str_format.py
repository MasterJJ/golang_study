age = 20
name ='swaRoop'

print '{} was {} years old when he worte this work'.format(age, name)

print 'Why is {}  playing with  that python '.format(name)



ifBoolean = False

if ifBoolean is True:
    print 'ifBoolean true'
else:
    print 'ifBoolean false'

isFullString = 'The Best is yet to come'

print isFullString

# add while


debugOut = 'current count is'

whileCount = 0
testStr = 'for test'
while whileCount < 10:
    whileCount += 1
    debugOut += testStr
    print debugOut


def funcTest():
    global debugOut
    print debugOut + 'call funcTest'

funcTest()




