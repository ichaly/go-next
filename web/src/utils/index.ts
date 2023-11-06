import {defu, createDefu} from 'defu'
import {extendTailwindMerge} from 'tailwind-merge'
import type {Strategy} from '~/types'

const customTwMerge = extendTailwindMerge({
    extend: {
        classGroups: {
            // icons: [(classPart: string) => /^i-/.test(classPart)]
        }
    }
})

const defuTwMerge = createDefu((obj, key, value, namespace) => {
    if (namespace !== 'default' && typeof obj[key] === 'string' && typeof value === 'string' && obj[key] && value) {
        // @ts-ignore
        obj[key] = customTwMerge(obj[key], value)
        return true
    }
})

export function mergeConfig<T>(strategy: Strategy, ...configs): T {
    if (strategy === 'override') {
        return defu({}, ...configs) as T
    }
    return defuTwMerge({}, ...configs) as T
}

export * from './lodash'