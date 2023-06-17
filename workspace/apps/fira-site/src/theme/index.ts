import { extendTheme, type ThemeConfig } from '@chakra-ui/react';

import { statTheme } from './components/stat.theme';

const colors = {
  primary: {
    50: '#d9fff6',
    100: '#acffe6',
    200: '#7bffd6',
    300: '#49ffc7',
    400: '#1affb7',
    500: '#00e69e',
    600: '#00b37b',
    700: '#008057',
    800: '#004e34',
    900: '#001c10',
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
