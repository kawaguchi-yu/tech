import { ChakraProvider } from "@chakra-ui/react"
import  Theme  from "../../public/theme"
function App({ Component, pageProps }) {
  return (
    <ChakraProvider theme={Theme}>
      <Component {...pageProps} />
    </ChakraProvider>
  )
}

export default App