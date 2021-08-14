import { ChakraProvider } from "@chakra-ui/react"
import  Theme  from "./components/theme"
function App({ Component, pageProps }) {
  return (
    <ChakraProvider theme={Theme}>
      <Component {...pageProps} />
    </ChakraProvider>
  )
}

export default App