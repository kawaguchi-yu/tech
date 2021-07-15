import React, { useState } from "react"
import { useForm } from "react-hook-form";
import Link from './components/Link';
import {
	chakra,
	Flex,
	Box,
	FormControl,
	FormLabel,
	Input,
	Checkbox,
	Stack,
	Button,
	Heading,
	useColorModeValue,
	FormHelperText,
} from '@chakra-ui/react';
type LoginData = {
	EMail: string;
	Password: string;
}
const userData: LoginData = {
	EMail: "",
	Password: "",
};
const Login = (): JSX.Element => {
	const { register, handleSubmit, formState, formState: { errors }, getValues } = useForm<LoginData>({
		mode: "onTouched"
	});
	const [posts, setPosts] = useState([]);

	const setData = () => {
		const hasData = getValues(["EMail", "Password"]);
		userData.EMail = hasData[0]
		userData.Password = hasData[1]
		console.log(userData)
	};

	const ApiFetch = () => {
		setData()
		fetch("http://localhost:8080/login", {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: JSON.stringify(userData),
		})
			.then((res) => res.json())
			.then((data) => {
				setPosts(data);})
			.catch((err)=>{console.log(err)})
			
	};

	return (
		<>
			<chakra.div>
				<Flex bg={useColorModeValue('gray.100', 'gray.900')}>
					<Link h={16} p={2} href="/">
						<Heading>Techer</Heading>
					</Link>
				</Flex>
				<Flex justify={'center'}>
					<Stack spacing={8} py={12}>
						<Stack align={'center'}>
							<Heading>Techerにログインする</Heading>
						</Stack>
						<Box
							bg={useColorModeValue('white', 'gray.700')}
							boxShadow={'lg'}
							p={8}>
							<Stack spacing={4}>

								<FormControl w={400} onSubmit={handleSubmit(setData)}
									isInvalid={errors.EMail ? true : false}>
									<Input
										type="email"
										placeholder="example@gmail.com"
										{...register("EMail", {
											required: "EMailを入力してください",
											pattern: {
												value: /^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/,
												message: 'メールアドレスを入力してください' // JS only: <p>error message</p> TS only support string
											},
										})}
									/>
									{errors.EMail && errors.EMail.message}
								</FormControl>

								<FormControl w={400} onSubmit={handleSubmit(setData)}
									isInvalid={errors.Password ? true : false}>
									<Input
										type="password"
										placeholder="password1"
										{...register("Password", {
											required: "パスワードを入力してください",
											minLength: {
												value: 8,
												message: '8文字以上にしてください' // JS only: <p>error message</p> TS only support string
											},
											pattern: {
												value: /^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d)[a-zA-Z\d]{8,100}$/,
												message: '小文字大文字数字をそれぞれ含めてください' // JS only: <p>error message</p> TS only support string
											}
										})}
									/>
									{errors.Password && errors.Password.message}
								</FormControl>

								<Stack align={`center`} spacing={10}>
									<Button type="submit"
										colorScheme="teal"
										onClick={ApiFetch}
										disabled={!formState.isValid}>
										ログイン
									</Button>
									<Heading>{JSON.stringify(posts)}</Heading>
								</Stack>
							</Stack>
						</Box>
					</Stack>
				</Flex>
			</chakra.div>
		</>
	);
}

export default Login