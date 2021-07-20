import React, { useState, useEffect } from "react"
import {
	chakra,
	Stack,
	Button,
	Heading,
} from '@chakra-ui/react';
import Link from './components/Link';
import MyPagesTemplate from "./myPagesTemplate";
type user = {
	ID: string
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Posts: string
	Profile: string
	ProfileID: string
	Goods: string
};
var userData: user = {
	ID: "",
	CreatedAt: "",
	UpdatedAt: "",
	DeletedAt: "",
	Name: "",
	EMail: "",
	Password: "",
	Posts: "",
	Profile: "",
	ProfileID: "",
	Goods: "",
}

const MyPages = (): JSX.Element => {
	useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
		})
			.then((res) => res.json())
			.then((data) => {
				const result = JSON.stringify(data)
				const result2: user = JSON.parse(result)
				userData = result2
				setEMail(userData)
				if (userData == null) {
					console.log("データはないよ！", userData)
				} else {
					setHasCookie(true)
					console.log("データはあるよ！", userData)
				}
			})
	}, [])
	const [email, setEMail] = useState<user>({ ID: "", CreatedAt: "", UpdatedAt: "", DeletedAt: "", Name: "", EMail: "", Password: "", Posts: "", Profile: "", ProfileID: "", Goods: "", });
	const [hasCookie, setHasCookie] = useState<boolean>(false);

	return (
		<>
			<MyPagesTemplate />
			<chakra.div>
				<Stack direction="row" align="center">
					{hasCookie
						? <><Heading>welcome!{email.Name}</Heading>
							<Heading><Link href="/config">設定する</Link></Heading>
						</>
						: <Heading>ログインしてないよ</Heading>}
				</Stack>
			</chakra.div>
		</>
	)
}
export default MyPages