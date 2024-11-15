import { ManagementPanel } from "./components/ManagementPanel";

export default function Home() {
	const user: User = null;
	return (
		<div>
			<ManagementPanel user={user} />
		</div>
	);
}
