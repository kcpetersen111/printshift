import { ManagementPanel } from "./components/ManagementPanel";
import { AccessLevel } from "./lib/enums/AccessLevel";
import { User } from "./lib/models/User";

export default function Home() {
	const user: User = {
		name: "Admin Dummy User",
		accessLevel: AccessLevel.Admin,
		email: "dummy@user.com",
		classes: [],
		printers: [],
		printersCanAssign: 1000
	};

	return (
		<div className="flex flex-col items-center">
			<ManagementPanel user={user} />
		</div>
	);
}
