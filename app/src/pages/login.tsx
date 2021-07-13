import React from "react"
import {
	Flex,
	Box,
	FormControl,
	FormLabel,
	Input,
	Checkbox,
	Stack,
	Link,
	Button,
	Heading,
	Text,
	useColorModeValue,
	FormHelperText,
} from '@chakra-ui/react';

const Login = (): JSX.Element => {
	return <>
		<Flex bg={useColorModeValue('gray.100', 'gray.900')} alignItems={'center'}>
			<Box h={16} p={2}>
				<Link href="/">
					<Heading>Techer</Heading>
				</Link>
			</Box>
		</Flex>
		<Flex
			minH={'100vh'}
			justify={'center'}
			bg={useColorModeValue('gray.50', 'gray.800')}>
			<Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6}>
				<Stack align={'center'}>
					<Heading fontSize={'4xl'}>Techerにログインする</Heading>
				</Stack>
				<Box
					rounded={'lg'}
					bg={useColorModeValue('white', 'gray.700')}
					boxShadow={'lg'}
					p={8}>
					<Stack spacing={4}>
						<FormControl id="email">
							<Input type="email"
								placeholder="メールアドレス" />
						</FormControl>
						<FormControl id="password">
							<Input type="password"
								placeholder="パスワード" />
						</FormControl>
						<Stack spacing={10}>
							<Stack
								direction={{ base: 'column', sm: 'row' }}
								align={'start'}
								justify={'space-between'}>
								<Checkbox>Remember me</Checkbox>
								<Link color={'blue.400'}>Forgot password?</Link>
							</Stack>
							<Button
								bg={'blue.400'}
								color={'white'}
								_hover={{
									bg: 'blue.500',
								}}>
								ログイン
							</Button>
						</Stack>
					</Stack>
				</Box>
			</Stack>
		</Flex>
	</>
}

export default Login