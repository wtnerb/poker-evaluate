# Evaluate
This microservice will evaluate and select the winner from a hand of poker.

The expected input is a poker table object as per the `poker/models` package. It will respond with the winner(s) of the hand.

Notably, how to handle ties or when someone has gone all-in add extra complexity to the problem. This will likely force the models to be updated.

Regardless, version 0.1 of this service is to recieve a post request with the table state information and (after validating that it is a game end table state) return with the winner of the hand.

## goals
## goals
- [x] Server accepts JSON post requests with table data
- [x] Can rank hands
- [x] Can build best hand
- [x] Can compare hands of same rank
- [ ] Can handle a tie
- [x] Containerized
- [ ] Deployed
- [ ] Security (does not hold information, but accessing with HTTPS secures information while in transit and verifying that a response came from this server will help as well)
- [ ] Keep test coverage over 90% while working

## Log
2019-03-19 v0.1 server works, can compare hands of different rank, some error checking, test coverage 91%  
2019-03-25 v0.2 can now evaluate all types of hands against all other types of hands, test coverage 90.5%