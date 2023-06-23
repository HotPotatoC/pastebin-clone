/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: '#0043FF',
        'primary-darken': '#003CE0',
        white: '#ECECEC',
        black: '#1B1B1B',
        grey: '#272727',
      },
      fontFamily: {
        'fira-code': ['"Fira Code"', 'monospace'],
      },
    },
  },
  plugins: [],
}
