import React, { ReactNode } from 'react';

interface PageTitleProps {
  children: ReactNode;
}

export default function PageTitle({ children }: PageTitleProps) {
  return <h1>{children}</h1>;
}
