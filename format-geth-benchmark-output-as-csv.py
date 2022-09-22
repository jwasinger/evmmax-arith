import sys, math
import re
import statistics

def insert_bench(benches_map, bench_preset, bench_op, bench_limbs, time):
    if not bench_op in benches_map:
        benches_map[bench_op] = {} 
    if not bench_preset in benches_map[bench_op]:
        benches_map[bench_op][bench_preset] = {}
    if not bench_limbs in benches_map[bench_op][bench_preset]:
        benches_map[bench_op][bench_preset][bench_limbs] = {'data':[]}

    benches_map[bench_op][bench_preset][bench_limbs]['data'].append(time)

#input_lines = sys.stdin.readlines()
input_lines = []
with open('non-unrolled_mulmont_bench.txt') as f:
    input_lines = f.readlines()

benches_map = {}

print("name, limb count, time (ns)")
for line in input_lines[4:-2]:
    parts = [elem for elem in line[13:].split(' ') if elem and elem != '\t']

    bench_full = parts[0][:-2]
    #bench_full = re.search(r'(.*)?(#.*-.*$)', parts[0]).groups()[0]
    if '#' in parts[0] and parts[0].index('#'):
        bench_full = parts[0].split('#')[0]

    bench_preset = bench_full.split('_')[0]
    bench_op = bench_full.split('_')[1]
    bench_limbs = bench_full.split('_')[2]
    time = float(parts[2])
    unit = re.search(r'(.*)\/', parts[3]).groups()[0]

    if unit != 'ns':
        raise Exception("expected ns got {}".format(unit))

    insert_bench(benches_map, bench_preset, bench_op, bench_limbs, time)

for bench_op in benches_map.keys():
    for bench_preset in benches_map[bench_op].keys():
        for bench_limbs in benches_map[bench_op][bench_preset].keys():
            item = benches_map[bench_op][bench_preset][bench_limbs]
            item['stddev'] = statistics.stdev(item['data'])
            item['mean'] = statistics.mean(item['data'])
            print("{},{},{},{},{}".format(bench_preset, bench_op, bench_limbs, item['mean'], item['stddev']))
