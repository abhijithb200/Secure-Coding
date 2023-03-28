# CodeGuardian

Description: CodeGuardian is a DevSecOps tool that helps to identify and remediate security vulnerabilities in source code. It is an automated tool that integrates with existing code repositories and performs static code analysis to detect security flaws and vulnerabilities.

Features:

Integration with popular code repositories like GitHub, Bitbucket, etc.
Automated static code analysis to detect security flaws and vulnerabilities
Integration with popular vulnerability databases like the National Vulnerability Database (NVD), Common Vulnerabilities and Exposures (CVE), etc.
Real-time security feedback and vulnerability reports to developers and security teams
Integration with popular issue tracking tools like JIRA, Trello, etc. to track and manage security issues
Customizable security policies and controls for code analysis
Integration with popular code editors like VSCode, Atom, etc.
How it works:

CodeGuardian integrates with the existing code repository used in the SDLC.
Automated static code analysis is performed to detect security flaws and vulnerabilities.
Real-time feedback and vulnerability reports are generated and sent to developers and security teams.
Security policies and controls can be customized for code analysis.
Integration with popular code editors allows developers to view and remediate vulnerabilities directly within their editor.
Benefits:

Early detection of security vulnerabilities in source code, reducing the risk of vulnerabilities and cyber attacks.
Improved collaboration between developers and security teams.
Time and cost savings due to automated code analysis.
Faster time to market, as security vulnerabilities are identified and remediated early in the development process.
Overall, CodeGuardian is a powerful tool for DevSecOps teams looking to automate code analysis and identify security vulnerabilities in source code.

# Vulnerability Detection Scope

## Reflected XSS

```php
echo "Name is :".$_GET['name'];
```

```php
$a = $_GET['name'];
echo "Name is".$a;
```

```php
$a = $_GET['name'];
$b = $a;
echo "Name is".$b;
```

```php
$a = "Name is ".$_GET['name'];
echo $a;
```

```php
$a = $_GET['name'];
$b = "Name is ".$a;
echo "I am ".$b;
```
