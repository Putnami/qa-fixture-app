import { describe, expect, it } from 'bun:test';
import { farewell } from '../src/farewell';

describe('farewell', () => {
  it('says goodbye to a person by name', () => {
    expect(farewell(' Ada ')).toBe('Goodbye, Ada!');
  });

  it('says goodbye to nobody in particular when the name is blank', () => {
    expect(farewell('')).toBe('Goodbye!');
  });
});
