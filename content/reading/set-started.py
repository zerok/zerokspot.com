import pathlib
def get_date(lines):
    for line in lines:
        elems = line.split(': ')
        if len(elems) > 1 and elems[0] == 'date':
            return elems[1]

for f in pathlib.Path('.').glob('*.md'):
    lines = f.read_text().split('\n')
    finished_idx = -1
    started_idx = -1
    for idx, l in enumerate(lines):
        if l.startswith('finished: '):
            finished_idx = idx
        elif l.startswith('started: '):
            started_idx = idx
    if started_idx > -1:
        continue
    if finished_idx == -1:
        continue
    lines.insert(finished_idx, 'started: {}'.format(get_date(lines)))
    print(f.name, started_idx, finished_idx)
    f.write_text('\n'.join(lines))

