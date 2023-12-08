#!/usr/bin/env ts-node
import * as fs from 'fs'

const data = JSON.parse(fs.readFileSync(`${__dirname}/../../../package.json`, 'utf8'))
const collection = [
  ...Object.keys(data.dependencies),
  ...Object.keys(data.devDependencies)
].filter(i => i.startsWith('@iconify-json/'))

for (let module of collection) {
  import(module).then(({ icons, info }) => {
    const keys = Object.keys(icons.icons)
    const tpl = `export default {
  name: '${info.name}',
  prefix: '${icons.prefix}',
  icons: [
    ${keys.map(i => `'i-${icons.prefix}:${i}',`).join('\n    ')}
  ],
}`
    fs.writeFile(`${__dirname}/data/icons.${icons.prefix}.ts`, tpl, (err) => {
      if (err) throw err
    })
  })
}


