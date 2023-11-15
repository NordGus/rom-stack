/** @type {import('tailwindcss').Config} */

export default {
  content: ["./html/templates/**/*.gohtml", "./client/**/*.{js,ts,jsx,tsx}"],
  theme: {
    darkMode: "class",
    fontFamily: {
      sans: ["Lato", "system-ui", "sans-serif"],
    },
    extend: {},
  },
  plugins: [],
};
