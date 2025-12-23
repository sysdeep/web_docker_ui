interface TotalReportProps {
  total: number;
  filtered: number;
}

export default function TotalReport({ total, filtered }: TotalReportProps) {
  return (
    <div>
      <p>
        Total: {filtered}/{total}
      </p>
    </div>
  );
}
