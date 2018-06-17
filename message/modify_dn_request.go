package message

import "fmt"

//
//        ModifyDNRequest ::= [APPLICATION 12] SEQUENCE {
//             entry           LDAPDN,
//             newrdn          RelativeLDAPDN,
//             deleteoldrdn    BOOLEAN,
//             newSuperior     [0] LDAPDN OPTIONAL }
func readModifyDNRequest(bytes *Bytes) (ret ModifyDNRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagModifyDNRequest, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readModifyDNRequest:\n%s", err.Error())}
		return
	}
	return
}
func (req *ModifyDNRequest) readComponents(bytes *Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	req.newrdn, err = readRelativeLDAPDN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	req.deleteoldrdn, err = readBOOLEAN(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	if bytes.HasMoreData() {
		var tag TagAndLength
		tag, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
			return
		}
		if tag.Tag == TagModifyDNRequestNewSuperior {
			var ldapdn LDAPDN
			ldapdn, err = readTaggedLDAPDN(bytes, classContextSpecific, TagModifyDNRequestNewSuperior)
			if err != nil {
				err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
				return
			}
			req.newSuperior = ldapdn.Pointer()
		}
	}
	return
}
