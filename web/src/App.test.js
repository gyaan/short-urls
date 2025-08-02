import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders sign in form', () => {
  const { getByText } = render(<App />);
  const signInElement = getByText('Sign in');
  expect(signInElement).toBeInTheDocument();
});
