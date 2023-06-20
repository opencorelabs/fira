import { extendTheme, type ThemeConfig } from '@chakra-ui/react';

import { statTheme } from './components/stat.theme';

const colors = {
  primary: {
    '50': '#EBF7F9',
    '100': '#C8E8EF',
    '200': '#A4D9E4',
    '300': '#81CADA',
    '400': '#5DBBD0',
    '500': '#3AACC5',
    '600': '#2E899E',
    '700': '#236776',
    '800': '#17454F',
    '900': '#0C2227',
  },
  secondary: {
    '50': '#F0F4F4',
    '100': '#D6E0E1',
    '200': '#BBCCCE',
    '300': '#A1B8BA',
    '400': '#86A4A7',
    '500': '#6C9093',
    '600': '#567476',
    '700': '#415758',
    '800': '#2B3A3B',
    '900': '#161D1D',
  },
  discord: {
    50: '#fefeff',
    100: '#ced2fb',
    200: '#b7bcf9',
    300: '#9fa6f8',
    400: '#8791f6',
    500: '#5865F2',
    600: '#2939ee',
    700: '#1225eb',
    800: '#1021d4',
    900: '#0d1aa4',
  },
};

const config: ThemeConfig = {
  initialColorMode: 'system',
  useSystemColorMode: true,
};

export const theme = extendTheme({ colors, config, components: { Stat: statTheme } });
