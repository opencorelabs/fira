import { Rasa } from 'next/font/google';
const rasa = Rasa({ weight: ['400', '700'], subsets: ['latin'] });

export function GlobalStyle() {
  return (
    <style jsx global>
      {`
        html {
          font-family: ${rasa.style.fontFamily};
        }
      `}
    </style>
  );
}
