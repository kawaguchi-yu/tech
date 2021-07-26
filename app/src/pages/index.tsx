import {
  chakra,
  Button,
  Stack,
} from "@chakra-ui/react"
import Link from './components/Link';
import React from "react"
import Template from "./template";
const Home = (): JSX.Element => {
  return (
    <>
      <chakra.div>
        <Template />
        クイズを投稿して知見を共有しませんか！
        <Stack direction="row" align="center">
          <Link href="/post">
            <Button colorScheme="teal" variant="solid">
              クイズを投稿する
            </Button>
          </Link>
        </Stack>
      </chakra.div>
    </>
  )
}

export default Home