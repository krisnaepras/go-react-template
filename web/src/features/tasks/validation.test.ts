import { describe, expect, it } from 'vitest'
import { validateTitle } from './validation'

describe('validateTitle', () => {
  it('rejects blank titles', () => expect(validateTitle('   ')).toBeTruthy())
  it('accepts ordinary titles', () => expect(validateTitle('Check inventory')).toBeNull())
  it('rejects overlong titles', () => expect(validateTitle('x'.repeat(201))).toBeTruthy())
})
