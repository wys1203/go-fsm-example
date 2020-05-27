## Introduction
FSM (finite-state machine). This is a simple example implement a FSM base on the features of go channel.
> https://stackoverflow.com/questions/38798863/golang-pause-a-loop-in-a-goroutine-with-channels/

### Receiving from nil channels:
>In Go, receiving from (or sending to) a nil channel results in "blocking forever". This in fact is a very important feature to implement the following trick: In a for-select pattern, if you set a case channel to nil, the corresponding case will not be matched in the next iteration. In other words, the case is "disabled".

### Receiving from closed channels:
>In Go, receiving from a closed channel always returns immediately. Therefore, you may replace your default case by a variable holding a closed channel. When the variable holds the closed channel, it behaves like the default case; However, when the variable holds nil, the case is never matched, having the "pause" behavior.
