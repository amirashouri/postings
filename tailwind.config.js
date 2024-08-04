/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'views/*.templ',
    'views/*/*.templ',
  ],
  daisyui: {
    themes: [
      {
        myTheme: {
          "primary": "#1b4f22",
          "secondary": "#48d7f7",
          "accent": "#37cdbe",
          "neutral": "#3d4451",
          "base-100": "#ffffff",
        },
      }, 
      // "dark",
      // "cupcake",
    ],
  },
  theme: {
    container: {
      center: true,
    },
    extend: {
      fontFamily: {
          sans: ['Graphik', 'sans-serif'],
          serif: ['Merriweather', 'serif'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('daisyui'),
  ],
}

