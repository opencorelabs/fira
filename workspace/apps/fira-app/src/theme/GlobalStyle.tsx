import { Roboto } from 'next/font/google';
const roboto = Roboto({ weight: ['400', '700'], subsets: ['latin'] });

export function GlobalStyle() {
  return (
    <style jsx global>
      {`
        html {
          font-family: ${roboto.style.fontFamily};
        }
      `}
    </style>
  );
}
