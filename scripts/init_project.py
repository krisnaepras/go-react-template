#!/usr/bin/env python3
"""Rename the starter once per new project. No external Python deps."""
from __future__ import annotations
import argparse
import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
DEFAULT_MODULE = 'example.com/office-template'

def main() -> None:
    parser = argparse.ArgumentParser(description='Initialize a project from the office starter template')
    parser.add_argument('--module', required=True, help='Go module path, e.g. github.com/acme/hr-platform')
    parser.add_argument('--title', required=True, help='Human-readable title, e.g. HR Platform')
    args = parser.parse_args()
    if not re.fullmatch(r'[a-zA-Z0-9._~/-]+', args.module) or '/' not in args.module:
        parser.error('--module must be a valid Go module path with a domain and path')
    mod = ROOT / 'go.mod'
    if f'module {DEFAULT_MODULE}' not in mod.read_text():
        parser.error('Starter already initialized: original module path not found')
    for file in [mod, *ROOT.rglob('*.go')]:
        # No need to modify generated code or module cache; starter contains none.
        file.write_text(file.read_text().replace(DEFAULT_MODULE, args.module))
    pkg = ROOT / 'web/package.json'
    obj = json.loads(pkg.read_text())
    slug = re.sub(r'[^a-z0-9-]+', '-', args.module.rsplit('/', 1)[-1].lower()).strip('-')
    obj['name'] = f'{slug}-web'
    pkg.write_text(json.dumps(obj, indent=2, ensure_ascii=False) + '\n')
    for rel in ('web/index.html', 'web/src/app/shell.tsx', 'web/src/features/dashboard/dashboard-page.tsx', 'api/openapi/openapi.yaml'):
        path = ROOT / rel
        txt = path.read_text().replace('Office Starter', args.title).replace('Office Go React Starter', args.title)
        path.write_text(txt)
    print(f'Initialized {args.title!r} ({args.module}). Run go mod tidy and npm install, then commit lockfiles.')

if __name__ == '__main__':
    main()
