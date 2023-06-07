Name:           lwr
Version:        0.1
Release:        1%{?dist}
Summary:        Lightweight replacement for LanSweeper Agent on Linux

License:        MIT
URL:            https://github.com/wanion/lwr
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang
BuildRequires:  make

Provides:       %{name} = %{version}

%description
Lightweight push-only client to replace LanSweeper Agent on Linux machines. It supports sending
reports directly to the LanSweeper scan server. Sending reports to cloud relay is not supported.

%global debug_package %{nil}

%prep
%autosetup


%build
make build


%install
rm -rf $RPM_BUILD_ROOT
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}

%files
%{_bindir}/%{name}


%changelog
* Wed Jun  9 2021 Fred Young <fred@wanion.net> - 0.1
- First release
