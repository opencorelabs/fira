import { statAnatomy } from '@chakra-ui/anatomy';
import { createMultiStyleConfigHelpers } from '@chakra-ui/react';

const { definePartsStyle, defineMultiStyleConfig } = createMultiStyleConfigHelpers(
  statAnatomy.keys
);

export const statTheme = defineMultiStyleConfig({
  variants: {
    networth: definePartsStyle({
      label: { fontSize: '2xl' },
      number: { fontSize: '2xl' },
      helpText: { fontSize: 'md' },
    }),
  },
  sizes: {
    sm: definePartsStyle({
      label: { fontSize: 'sm' },
      helpText: { fontSize: 'sm' },
      number: { fontSize: 'sm' },
    }),
    xl: definePartsStyle({
      label: { fontSize: 'xl' },
      helpText: { fontSize: 'xl' },
      number: { fontSize: 'xl' },
    }),
  },
});
