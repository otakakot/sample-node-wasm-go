import "./wasm_exec.js";

import { readFileSync } from 'node:fs';

const go = new Go();

const source = readFileSync('main.wasm')

const { instance } = await WebAssembly.instantiate(source, go.importObject);

go.run(instance);

golog("Hello, World!");
