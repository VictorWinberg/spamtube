// Styles
import "@mdi/font/css/materialdesignicons.css";
import "vuetify/styles";

// Vuetify
import { createVuetify, ThemeDefinition } from "vuetify";

const brandTheme: ThemeDefinition = {
  dark: false,
  colors: {
    background: "#000000",
    backgroundFade: "#3a668b",
    primary: "#58b3de",
    secondary: "304362",
    tetriary: "#2b4062",
    darkText: "#000000",
    lightText: "#ffffff",
    button: "#6eb2da",
    buttonText: "#000000",
    info: "##0074D9",
    success: "##2ECC40",
    warning: "##FFDC00",
    error: "#FF4136",
  },
};

export default createVuetify({
  theme: {
    defaultTheme: "brandTheme",
    themes: {
      brandTheme,
    },
  },
});
// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
