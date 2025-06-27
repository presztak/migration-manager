import { useNavigate, useParams } from "react-router";
import TabView from "components/TabView";
import NetworkOverview from "pages/NetworkOverview";

const NetworkDetail = () => {
  const { name, activeTab } = useParams<{ name: string; activeTab: string }>();
  const navigate = useNavigate();

  const tabs = [
    {
      key: "overview",
      title: "Overview",
      content: <NetworkOverview />,
    },
  ];

  return (
    <div className="d-flex flex-column">
      <div className="scroll-container flex-grow-1 p-3">
        <TabView
          defaultTab="overview"
          activeTab={activeTab}
          tabs={tabs}
          onSelect={(key) => navigate(`/ui/networks/${name}/${key}`)}
        />
      </div>
    </div>
  );
};

export default NetworkDetail;
