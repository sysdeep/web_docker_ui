interface TotalReportProps {
  total: number;
}

export default function TotalReport({ total }: TotalReportProps) {
  return (
    <div>
      <p>Total: {total}</p>
    </div>
  );
}
