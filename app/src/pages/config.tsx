import React, { useState } from "react"
import {
	Box,
	Flex,
	Button,
	Input,
	Image,
} from '@chakra-ui/react';
import MyPagesTemplate from "./myPagesTemplate";
const Config = () => {
	const [view, setview] = useState<string>();
	const [posts, setPosts] = useState<Blob>();
	const [iconData, setIconData] = useState<FormData>();
	const onFileInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setview(window.URL.createObjectURL(e.target.files[0]))
		const image = new FormData()
		image.append("file", e.target.files[0])
		setIconData(image)
		console.log(e.target.files[0])
	};
	const ApiFetch = () => {
		const options: RequestInit = {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "multipart/form-data", }, // JSON形式のデータのヘッダー
			credentials: 'include',
			body: iconData,
		}
		delete options.headers["Content-Type"];
		fetch("http://localhost:8080/seticon", options)
			.then((res) => res.blob())
			.then((data) => {
				setPosts(data);
				console.log("返ってきたデータ")
				console.log(data)
				console.log(view)
			})
			// .catch((err) => { console.log(err) })
		console.log("アイコンデータ",iconData)
	};
	return (<>
		<MyPagesTemplate />
		<Flex>
			{view &&
				<Image boxSize="300px" src={view} alt="select picture" />
			}
			<Flex>
				<Input name="file" type='file' accept="image/*" onChange={onFileInputChange} />
			</Flex>
		</Flex>
		<Flex>
			<Button onClick={ApiFetch}>アイコンを変更する</Button>
			{posts &&
				<Image boxSize="300px" src={(window.URL.createObjectURL(posts))} alt="select picture" />
			}
		</Flex>
	</>)
}
export default Config