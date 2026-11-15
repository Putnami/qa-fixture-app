import { describe, expect } from 'bun:test';
import { specTest } from '@putnami/runtime/spectest';
import { greet } from '../src/greeting';

const GREETS = { feature: 'greeter/greeting', requirement: 'greets-by-name', check: 'greeter-greeting' };
const CAPITALIZES = { feature: 'greeter/greeting', requirement: 'capitalizes-the-name', check: 'greeter-greeting' };

describe('greet', () => {
  specTest('greets a person by name', GREETS, () => {
    expect(greet('Ada')).toBe('Hello, Ada!');
  });

  specTest('greets nobody in particular when the name is blank', GREETS, () => {
    expect(greet('  ')).toBe('Hello!');
  });

  specTest('capitalizes the first letter of the name', CAPITALIZES, () => {
    expect(greet('grace')).toBe('Hello, Grace!');
  });
});
