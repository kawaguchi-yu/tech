import {
	Box,
	Flex,
	Button,
	useColorModeValue,
	Spacer,
	Heading,
} from '@chakra-ui/react';
import Link from './components/Link';
 const Common = () => {

	return (
		<>
			<Flex bg={useColorModeValue('gray.100', 'gray.900')}alignItems={'center'}>
				<Box h={16} p={2}>
				<Link href="/">
					<Heading>Techer</Heading>
					</Link>
				</Box>
				<Spacer />
				<Box mr={4}>
				<Button mr="4" colorScheme="teal" variant="solid">
					<Link href="/Registration">
						ユーザー登録
					</Link>
				</Button>
				<Button colorScheme="teal">
					<Link href="/login">
						ログイン
					</Link>
				</Button>
				</Box>
			</Flex>
		</>
	);
}
export default Common