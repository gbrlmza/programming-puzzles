#!/usr/bin/python3

def search(engines, queries):
    switches = 0

    while True:
        numq = 0 # number of queries that a given engine can handle until a switching is required
        
        for e in engines:
            # chose the engine that can process the max number of queries
            if e not in queries: # the engine can process all the remaining queries
                return switches
            numq = max(queries.index(e), numq)
        
        queries = queries[numq:] # remove processed queries
        
        if not len(queries) or not numq:
            break

        switches += 1
    
    return switches


def main():
    cases = int(input())  # read a line with a single integer
    
    for case in range(1, cases+1):
        items = int(input())
        engines = [input() for x in range(items)]
        items = int(input())
        queries = [input() for x in range(items)]

        switches = search(engines, queries)
        print('Case #{}: {}'.format(case, switches))

main()
