#!/usr/bin/python3
import sys

def get_trains(tatime, abtrips, batrips):
    start = {'A':0, 'B':0}

    # Unify and sort by departure time all trips
    trips = list(map(lambda x: x + ['A'], abtrips)) + list(map(lambda x: x + ['B'], batrips))
    trips.sort(key=lambda t: t[0])

    # Calculate train disponibility for each station for arrivals
    ready = {}
    ready['A'] = [x[1] + tatime for x in trips if x[2] == 'B']
    ready['B'] = [x[1] + tatime for x in trips if x[2] == 'A']
    
    for t in trips:
        train_available = False
        
        for i, r in enumerate(ready[t[2]]):
            if r <= t[0]:
                train_available = True
                del(ready[t[2]][i])
                break

        if not train_available:
            start[t[2]] += 1

    return start

def main():
    cases = int(input())

    for case in range(1, cases+1):
        tatime = int(input())
        nabtrips, nbatrips = [int(x) for x in input().split(" ")]

        # Convert time from HH:MM to minutes to simplify posterior calculations
        abtrips = [[int(y.split(":")[0]) * 60 + int(y.split(":")[1]) for y in input().split(" ")] for x in range(nabtrips)]
        batrips = [[int(y.split(":")[0]) * 60 + int(y.split(":")[1]) for y in input().split(" ")] for x in range(nbatrips)]

        starting_trains = get_trains(tatime, abtrips, batrips)
        print("Case #{}: {} {}".format(case, starting_trains['A'], starting_trains['B']))

main()
