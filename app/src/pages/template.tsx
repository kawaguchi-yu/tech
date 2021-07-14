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
			<Flex bg={useColorModeValue('gray.100', 'gray.900')} alignItems={'center'}>
				<Link href="/">
					<Box h={16} p={2} color="Highlight">
						<Heading>Techer</Heading>
					</Box>
				</Link>
				<Spacer />
				<Box mr={4}>
					<Link href="/Registration">
						<Button mr="4" colorScheme="teal" variant="solid">
							ユーザー登録
						</Button>
					</Link>
					<Link href="/login">
						<Button colorScheme="teal">
							ログイン
						</Button>
					</Link>
				</Box>
			</Flex>
		</>
	);
}
export default Common