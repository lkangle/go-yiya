const extendSpacing = ({theme}) => ({
  ...theme('spacing'),
});

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      spacing: new Array(600).fill(0).reduce((r,_,index) => ({
        ...r, [index]: `${index}px`
      }), {}),
      fontSize: extendSpacing,
      lineHeight: extendSpacing,
    },
  },
  plugins: [],
}
