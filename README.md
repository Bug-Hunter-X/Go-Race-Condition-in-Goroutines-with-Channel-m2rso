# Go Race Condition in Goroutines with Channels

This repository demonstrates a race condition in a Go program that uses goroutines and channels. The program creates multiple goroutines that attempt to receive from a channel that only has one value sent to it. This can lead to unexpected behavior or deadlock.