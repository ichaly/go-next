import sampleSize from 'lodash/sampleSize'

/**
 * 获取任意长度的随机数字字母组合字符串
 * @param size 随机字符串的长度
 */
export function randomString(size: number): string {
  const charSet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  return sampleSize(charSet, size).toString().replace(/,/g, '')
}